var window = self;

function resolve(p, v = window) {
  parts = p.split(".");
  while (parts.length > 0) {
    v = v[parts.shift()];
  }
  return v;
}

const routes = new Map(
  JSON.parse(`{{jsonMarshal .Routes}}`).map(({ path, handler }) => [
    path,
    handler,
  ])
);

self.importScripts(...JSON.parse(`{{jsonMarshal .Scripts}}`));

self.addEventListener("error", (error) => {
  console.error("error", error);
});

self.addEventListener("install", (event) => {
  // Activate immediately
  self.skipWaiting();
});

self.addEventListener("activate", (event) => {
  // Take control of uncontrolled clients as soon as this worker activates.
  event.waitUntil(self.clients.claim());
});

self.addEventListener("fetch", (event) => {
  const req = event.request;
  const url = new URL(req.url);

  if (
    req.method === "GET" &&
    url.origin === self.location.origin &&
    routes.has(url.pathname)
  ) {
    event.respondWith(
      (async () => {
        try {
          fn = resolve(routes.get(url.pathname), self);
          return await fn(req);
        } catch (err) {
          const msg = "SW fetch handler error: " + err.message;
          console.error(msg);
          return new Response(msg, {
            status: 500,
            statusText: "internal server error",
          });
        }
      })()
    );
    return;
  }

  // Default: just forward the request to network
  event.respondWith(fetch(event.request));
});
