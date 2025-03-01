let isZoomOut = false;

function imageShow(url = "") {
    console.log("Image clicked");

    const zoomImage = document.getElementById("zoomImage");
    const zoomImageSrc = document.getElementById("zoomImageSrc");

    if (isZoomOut) {
        document.body.style.overflow = "auto";
        zoomImage.style.display = "none";
        isZoomOut = false;
        return;
    }

    if (url) {
        zoomImageSrc.src = url;
    }
    document.body.style.overflow = "hidden";
    zoomImage.style.display = "flex";
    isZoomOut = true;
}