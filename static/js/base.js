// Função para mostrar notificações
function showNotification(message, type = 'success') {
    const notification = createNotificationElement(message, type);
    document.getElementById('notifications').appendChild(notification);

    setTimeout(() => {
        notification.remove();
    }, 3000);
}

// Função auxiliar para criar o elemento da notificação
function createNotificationElement(message, type) {
    const notification = document.createElement('div');
    notification.className = `notification ${type}`;

    const icon = document.createElement('span');
    icon.className = 'icon';
    icon.innerHTML = getNotificationIcon(type);

    notification.appendChild(icon);
    notification.appendChild(document.createTextNode(message));

    return notification;
}

// Função para obter o ícone com base no tipo de notificação
function getNotificationIcon(type) {
    switch (type) {
        case 'success': return '✅';
        case 'error': return '❌';
        default: return '⚠️';
    }
}

// Função para alternar o estado da sidebar (colapsada ou expandida)
function toggleSidebar() {
    const sidebar = document.getElementById('sidebar');
    const texts = document.querySelectorAll('.sidebar-text');

    const isMobile = window.innerWidth <= 768;

    if (isMobile) {
        toggleSidebarForMobile(sidebar, texts);
    } else {
        toggleSidebarForDesktop(sidebar, texts);
    }
}

// Função para alternar o estado da sidebar em dispositivos móveis
function toggleSidebarForMobile(sidebar, texts) {
    if (sidebar.classList.contains('sidebar-collapsed')) {
        sidebar.classList.remove('sidebar-collapsed');
        sidebar.classList.add('sidebar-expanded');
        texts.forEach(text => text.classList.remove('hidden'));
        localStorage.setItem('sidebarState', 'expanded');
    } else {
        sidebar.classList.add('sidebar-collapsed');
        sidebar.classList.remove('sidebar-expanded');
        texts.forEach(text => text.classList.add('hidden'));
        localStorage.setItem('sidebarState', 'collapsed');
    }
}

// Função para alternar o estado da sidebar em dispositivos grandes
function toggleSidebarForDesktop(sidebar, texts) {
    if (sidebar.classList.contains('sidebar-collapsed')) {
        sidebar.classList.remove('sidebar-collapsed');
        sidebar.classList.add('sidebar-expanded');
        texts.forEach(text => text.classList.remove('hidden'));
        localStorage.setItem('sidebarState', 'expanded');
    } else {
        sidebar.classList.add('sidebar-collapsed');
        sidebar.classList.remove('sidebar-expanded');
        texts.forEach(text => text.classList.add('hidden'));
        localStorage.setItem('sidebarState', 'collapsed');
    }
}

// Função para inicializar o estado da sidebar ao carregar a página
function initializeSidebarState() {
    const sidebar = document.getElementById('sidebar');
    const texts = document.querySelectorAll('.sidebar-text');
    const sidebarState = localStorage.getItem('sidebarState');
    const isMobile = window.innerWidth <= 768;

    if (sidebar) {
        if (isMobile) {
            setSidebarForMobile(sidebar, texts, sidebarState);
        } else {
            setSidebarForDesktop(sidebar, texts, sidebarState);
        }
    }
}

// Função para definir o estado da sidebar em dispositivos móveis
function setSidebarForMobile(sidebar, texts, sidebarState) {
    if (sidebarState === 'collapsed') {
        sidebar.classList.add('sidebar-collapsed');
        sidebar.classList.remove('sidebar-expanded');
        texts.forEach(text => text.classList.add('hidden'));
    } else {
        sidebar.classList.add('sidebar-expanded');
        sidebar.classList.remove('sidebar-collapsed');
        texts.forEach(text => text.classList.remove('hidden'));
    }
}

// Função para definir o estado da sidebar em dispositivos grandes
function setSidebarForDesktop(sidebar, texts, sidebarState) {
    if (sidebarState === 'collapsed') {
        sidebar.classList.add('sidebar-collapsed');
        sidebar.classList.remove('sidebar-expanded');
        texts.forEach(text => text.classList.add('hidden'));
    } else {
        sidebar.classList.add('sidebar-expanded');
        sidebar.classList.remove('sidebar-collapsed');
        texts.forEach(text => text.classList.remove('hidden'));
    }
}

// Função para tratar a resposta do HTMX e exibir notificações ou redirecionamento
function handleHTMXResponse(event) {
    const response = event.detail.xhr.response;

    try {
        const parsedResponse = typeof response === 'string' ? JSON.parse(response) : response;

        if (parsedResponse) {
            if (!parsedResponse.redirect && parsedResponse.message) {
                showNotification(parsedResponse.message, parsedResponse.type);
                return;
            }

            if (parsedResponse.message) {
                localStorage.setItem('notification', JSON.stringify({
                    message: parsedResponse.message,
                    type: parsedResponse.type
                }));
            }

            if (parsedResponse.redirect) {
                const contentWrapper = document.getElementById('content-wrapper');
                contentWrapper.classList.remove('opacity-100');
                contentWrapper.classList.add('opacity-0');

                setTimeout(() => {
                    window.location.href = parsedResponse.redirect;
                }, 300);
            }
        }
    } catch (error) {
        console.log("Erro ao processar a resposta:", error);
    }
}

// Inicializa o script ao carregar o DOM
window.addEventListener('DOMContentLoaded', () => {
    initializeSidebarState();

    // Exibe a notificação armazenada, se houver
    const storedNotification = localStorage.getItem('notification');
    if (storedNotification) {
        const { message, type } = JSON.parse(storedNotification);
        showNotification(message, type);
        localStorage.removeItem('notification');
    }

    // Associa o evento do HTMX
    document.body.addEventListener('htmx:afterSwap', handleHTMXResponse);
});
