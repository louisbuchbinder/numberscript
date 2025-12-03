const serviceWorkerMainInit: Promise<boolean> =
  (async function (): Promise<boolean> {
    if ("serviceWorker" in navigator) {
      return navigator.serviceWorker
        .register("/sw.js", { scope: "/", updateViaCache: "imports" })
        .then((reg) => reg.update())
        .then(() => {
          if (window.navigator.serviceWorker.controller === null) {
            window.location.reload();
          }
          return true;
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
