self.addEventListener('push', function(event) {
    const options = {
        body: event.data.text(),
    };
    event.waitUntil(
        self.registration.showNotification('Notification Title', options)
    );
});
