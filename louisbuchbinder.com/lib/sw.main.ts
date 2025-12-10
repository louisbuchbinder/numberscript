const serviceWorkerMainInit: Promise<ServiceWorkerRegistration> =
  (async function (): Promise<ServiceWorkerRegistration> {
    if ("serviceWorker" in navigator) {
      return navigator.serviceWorker
        .register("/sw.js", { scope: "/", updateViaCache: "imports" })
        .then(async (reg) => {
          await reg.update();
          if (window.navigator.serviceWorker.controller === null) {
            window.location.reload();
            return;
          }
          return navigator.serviceWorker.ready;
        });
    } else {
      const err = new Error(
        "navigator.serviceWorker not supported in this browser."
      );
      const e = document.getElementById("page-error");
      if (Boolean(e)) {
        e.querySelector("p").textContent = err.message;
        e.removeAttribute("hidden");
      }
      throw err;
    }
  })();
