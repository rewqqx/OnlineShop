package selenium.utils

import config.Config
import org.opencv.core.*
import org.opencv.imgcodecs.Imgcodecs
import org.opencv.imgproc.Imgproc

private var counter = 0

fun Rect.getSquare(): Int {
    return this.width * this.height
}

fun Rect.equal(rect: Rect): Boolean {
    if (rect.x != x) {
        return false
    }

    if (rect.y != y) {
        return false
    }

    if (rect.width != width) {
        return false
    }

    if (rect.height != height) {
        return false
    }

    return true
}

fun Rect.getPoint(): org.openqa.selenium.Point {
    return org.openqa.selenium.Point(this.x, this.y)
}

fun Rect.getCenterPoint(): org.openqa.selenium.Point {
    return org.openqa.selenium.Point(this.x + this.width / 2, this.y + this.height / 2)
}

fun Rect.getOffsetPoint(x: Int, y: Int): org.openqa.selenium.Point {
    return org.openqa.selenium.Point(this.x + x, this.y + y)
}

fun Rect.getCenterPointOffset(x: Int, y: Int): org.openqa.selenium.Point {
    return org.openqa.selenium.Point(this.x + this.width / 2 + x, this.y + this.height / 2 + y)
}

fun Mat.outlineBinaryImage(radius: Int): Mat {
    val emptyImage = Mat.zeros(this.size(), this.type())

    for (i in 0 until this.rows()) {
        for (j in 0 until this.cols()) {

            // Check if the pixel is white
            if (this.get(i, j)[0] > 0) {

                // Set the corresponding pixel in the new empty image to white with radius 10
                Imgproc.circle(emptyImage, Point(j.toDouble(), i.toDouble()), radius, Scalar(255.0), -1)
            }
        }
    }

    return emptyImage
}

fun Mat.findDifference(mat: Mat, ignore: Rect?): Rect {
    val difference = Mat()
    Core.absdiff(this, mat, difference)

    val gray = Mat()
    Imgproc.cvtColor(difference, gray, Imgproc.COLOR_BGR2GRAY)

    // Threshold the grayscale image to create a binary image
    val binary = Mat()
    Imgproc.threshold(gray, binary, 0.0, 255.0, Imgproc.THRESH_BINARY or Imgproc.THRESH_OTSU)

    val outlined = binary.outlineBinaryImage(5);

    saveImageDebug(mat)
    saveImageDebug(this)
    saveImageDebug(binary)
    saveImageDebug(outlined)

    // Find the contours in the binary image
    val contours = mutableListOf<MatOfPoint>()
    val hierarchy = Mat()
    Imgproc.findContours(
        outlined,
        contours,
        hierarchy,
        Imgproc.RETR_EXTERNAL,
        Imgproc.CHAIN_APPROX_SIMPLE
    )

    // Find the bounding rectangles of the contours
    var boundingRects = mutableListOf<Rect>()
    for (contour in contours) {
        val rect = Imgproc.boundingRect(contour)
        boundingRects.add(rect)
    }

    if (ignore != null) {
        boundingRects = boundingRects.filter { rect -> !rect.equal(ignore) }.toMutableList()
    }


    val sortedBoundingRects = boundingRects.sortedBy { rect -> -rect.getSquare() }

    return if (sortedBoundingRects.isNotEmpty()) sortedBoundingRects[0] else Rect()
}

fun saveImageDebug(mat: Mat) {
    if (!Config.DEBUG) {
        return
    }

    Imgcodecs.imwrite("./output/image-" + counter + "-debug.jpg", mat)
    counter += 1
}