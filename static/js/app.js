(function() {
    /**@type{HTMLCanvasElement | null}*/
    const canvas = document.getElementById("drawingCanvas");
    if (canvas === null) {
        console.error("Could not find element with id: drawingCanvas");
        return;
    }

    const ctx = canvas.getContext("2d");

    /**@param{ HTMLCanvasElement | null } canvas*/
    function matchCanvasDimensionToWindow(canvas) {
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

    window.addEventListener("DOMContentLoaded", matchCanvasDimensionToWindow({ canvas }));
    window.addEventListener("resize", matchCanvasDimensionToWindow({ canvas }));

    canvas.addEventListener("pointerdown", startDrawing)
    canvas.addEventListener("pointermove", draw)
    canvas.addEventListener("pointerup", stopDrawing)

    /**@param{PointerEvent} pointerEvent */
    function getPositions(pointerEvent) {
    }

    function draw() {

    }

    function startDrawing() {

    }

    function stopDrawing() {

    }

})();
