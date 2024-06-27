(function() {
    function matchCanvasDimensionToWindow() {
        /**@type{ HTMLCanvasElement | null }*/
        const canvas = document.getElementById("drawingCanvas");
        if (canvas === null) {
            console.error("Could not find element with id: drawingCanvas");
            return;
        }

        const timeout = setTimeout(function() {
            clearTimeout(timeout)
            canvas.width = window.innerWidth
            canvas.height = window.innerHeight
        }, 300)
    }

    window.addEventListener("DOMContentLoaded", matchCanvasDimensionToWindow);
    window.addEventListener("resize", matchCanvasDimensionToWindow);
})();
