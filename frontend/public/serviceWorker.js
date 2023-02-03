const cacheName = 'localstalley_cache'
const contentToCache = [
	'./favicon.ico',
	'./index.html',
	'./logo192.png',
	'./logo512.png',
	'./robots.txt'
]

//self.addEventListener("install", (e) => {
  //console.log("[Service Worker] Install");
  //e.waitUntil(
    //(async () => {
      //const cache = await caches.open(cacheName);
      //console.log("[Service Worker] Caching all: app shell and content");
      //await cache.addAll(contentToCache);
    //})()
  //);
//});

self.addEventListener("install", event => {
    self.registration.showNotification('You are all set to receive Notification', {
      body: "Promise, We won't spam",
    })
  console.log("Service worker installed");
});

self.addEventListener("activate", event => {
   console.log("Service worker activated");
});

self.addEventListener("fetch", (e) => {
	return
  //console.log(`Service worker Fetched resource ${e.request.url}`);
});

self.addEventListener('push', function(event) {
  // Retrieve the textual payload from event.data (a PushMessageData object).
  // Other formats are supported (ArrayBuffer, Blob, JSON), check out the documentation
  // on https://developer.mozilla.org/en-US/docs/Web/API/PushMessageData.
  const payload = event.data ? event.data.text() : 'no payload';

  const {title, content} = JSON.parse(payload)
  // Keep the service worker alive until the notification is created.
  event.waitUntil(
    // Show a notification with title 'ServiceWorker Cookbook' and use the payload
    // as the body.
    self.registration.showNotification(title, {
      body: content,
    })
  );
});


