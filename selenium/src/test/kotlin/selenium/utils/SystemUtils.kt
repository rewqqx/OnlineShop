package selenium.utils

import io.kotlintest.shouldNotBe
import java.io.File
import java.nio.file.Paths
import java.util.concurrent.TimeUnit

val path = Paths.get("").toAbsolutePath().toString()
val downloadDir = path + "/src/test/resources/tmp".replace("/", File.separator)

fun waitForFileDownload(): String {
    var counter = 0

    while (File(downloadDir).listFiles()?.any { it.name.endsWith(".html") } == false) {
        TimeUnit.SECONDS.sleep(1)
        counter += 1
        counter.shouldNotBe(10)
    }

    return File(downloadDir).listFiles().find { it.name.endsWith(".html") }!!.path
}

fun clearDownloadedFiles() {
    val files = File(downloadDir).listFiles()
    for (file in files) {
        file.delete()
    }
}