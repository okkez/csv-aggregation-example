plugins {
    kotlin("jvm") version "1.4.21"
}

group = "net.okkez"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    implementation("org.apache.commons:commons-csv:1.8")
}

val fatJar = task("fatJar", type = Jar::class) {
    manifest {
        attributes["Main-Class"] = "MainKt"
    }

    from(
        configurations.runtimeClasspath.get().map {
            if (it.isDirectory()) it else zipTree(it)
        }
    )
    with(tasks.jar.get() as CopySpec)
}