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
    }, 2000);
}

document.body.addEventListener('htmx:afterSwap', (event) => {
    const response = event.detail.xhr.response;

    try {
        const parsedResponse = typeof response === 'string' ? JSON.parse(response) : response;


        if (parsedResponse && parsedResponse.message) {
            showNotification(parsedResponse.message, parsedResponse.type);

            if (parsedResponse.redirect) {
                setTimeout(() => {
                    window.location.href = parsedResponse.redirect;
                }, 2000);
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
    const texts = document.querySelectorAll('.sidebar-text');
    const sidebarState = localStorage.getItem('sidebarState');

    if (sidebarState === 'collapsed') {
        sidebar.classList.add('sidebar-collapsed');
        texts.forEach(text => text.classList.add('hidden'));
    } else {
        sidebar.classList.add('sidebar-expanded');
        texts.forEach(text => text.classList.remove('hidden'));
    }
});