function showNotification(message, type = 'success') {
    const notification = document.createElement('div');
    notification.className = `notification ${type}`;

    const icon = document.createElement('span');
    icon.className = 'icon';
    icon.innerHTML = type === 'success' ? '✅' : '❌';

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




