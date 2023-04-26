package opencv.start
import org.junit.jupiter.api.Test
import org.opencv.core.Mat


class OpenCVStartTest {

    init {
        nu.pattern.OpenCV.loadShared()
        System.loadLibrary(org.opencv.core.Core.NATIVE_LIBRARY_NAME)
    }

    @Test
    fun screenshotJSONTest() {
        val mat = Mat()
    }

}