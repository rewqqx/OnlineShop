plugins {
    kotlin("jvm") version "1.8.0"
    application
}

group = "org.example"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}


dependencies {
    // Tests Implementation
    testImplementation(kotlin("test"))
    implementation("io.kotlintest:kotlintest-runner-junit5:3.1.8")

    // Selenium Implementation
    implementation("org.seleniumhq.selenium:selenium-chrome-driver:4.8.1")
    implementation("org.seleniumhq.selenium:selenium-java:4.8.1")

    // OpenCV Implementation
    implementation("org.openpnp:opencv:4.5.1-2")
    implementation("org.apache.commons:commons-math3:3.6.1")
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(11)
}

application {
    mainClass.set("MainKt")
}

configurations {
    testCompileClasspath {
        attributes {
            attribute(TargetJvmVersion.TARGET_JVM_VERSION_ATTRIBUTE, 11)
        }
    }
}