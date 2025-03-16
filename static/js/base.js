function showNotification(message, type = 'success') {
    const notification = document.createElement('div');
    notification.className = `notification ${type}`;

    const icon = document.createElement('span');
    icon.className = 'icon';

    switch (type) {
        case 'success':
            icon.innerHTML = '✅';
            break;
        case 'error':
            icon.innerHTML = '❌';
            break;
        default:
            icon.innerHTML = '⚠️';
    } 

    notification.appendChild(icon);
    notification.appendChild(document.createTextNode(message));

    document.getElementById('notifications').appendChild(notification);

    setTimeout(() => {
        notification.remove();
    }, 3000);
}

document.body.addEventListener('htmx:afterSwap', (event) => {
    const response = event.detail.xhr.response;

    try {
        const parsedResponse = typeof response === 'string' ? JSON.parse(response) : response;

        if (parsedResponse) {
            if (!parsedResponse.redirect && parsedResponse.message) {
                showNotification(parsedResponse.message, parsedResponse.type);
                return
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
});


function toggleSidebar() {
    const sidebar = document.getElementById('sidebar');
    const texts = document.querySelectorAll('.sidebar-text');

    if (sidebar.classList.contains('sidebar-expanded')) {
        sidebar.classList.remove('sidebar-expanded');
        sidebar.classList.add('sidebar-collapsed');
        texts.forEach(text => text.classList.add('hidden'));

        localStorage.setItem('sidebarState', 'collapsed');
    } else {
        sidebar.classList.remove('sidebar-collapsed');
        sidebar.classList.add('sidebar-expanded');
        texts.forEach(text => text.classList.remove('hidden'));

        localStorage.setItem('sidebarState', 'expanded');
    }
}

window.addEventListener('DOMContentLoaded', () => {
    const sidebar = document.getElementById('sidebar');
    
    if (sidebar) {
        const texts = document.querySelectorAll('.sidebar-text');
        const sidebarState = localStorage.getItem('sidebarState');

        if (!sidebarState) {
            sidebar.classList.add('sidebar-expanded');
            texts.forEach(text => text.classList.remove('hidden'));
            localStorage.setItem('sidebarState', 'expanded');  // Grava o estado padrão
        } else if (sidebarState === 'collapsed') {
            sidebar.classList.add('sidebar-collapsed');
            texts.forEach(text => text.classList.add('hidden'));
        } else {
            sidebar.classList.add('sidebar-expanded');
            texts.forEach(text => text.classList.remove('hidden'));
        }
    }

    const storedNotification = localStorage.getItem('notification');

    if (storedNotification) {
        const { message, type } = JSON.parse(storedNotification);
        showNotification(message, type);

        localStorage.removeItem('notification');
    }
});