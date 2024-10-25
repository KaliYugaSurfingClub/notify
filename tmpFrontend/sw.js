self.addEventListener('push', event => {
    console.log('[Service Worker] Push Received.');
    console.log(`[Service Worker] Push had this data: "${event.data.text()}"`);

    const title = 'Test Webpush';
    const options = {
        body: event.data.text(),
        icon: 'icon.png',
        badge: 'badge.png',
        actions: [
            {action: 'view', title: 'Посмотреть'},
            {action: 'dismiss', title: 'Удалить'}
        ],
        data: {
            url: 'https://example.com' // URL to open on notification click
        }
    };

    event.waitUntil(
        self.registration.showNotification(title, options)
            .then(() => console.log('Notification displayed'))
            .catch(error => console.error('Error displaying notification:', error))
    );
});