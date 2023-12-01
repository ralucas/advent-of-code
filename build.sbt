organization := "advent-of-code"
version := "1.0.0"
scalaVersion := "3.2.2"

lazy val root = (project in file("."))
  .settings(
    // set the name of the project
    name := "Advent of Code",

    ivyLoggingLevel := UpdateLogging.Full,

    libraryDependencies += "com.typesafe" % "config" % "1.4.2"

    // only show warnings and errors on the screen for compilations.
    //  this applies to both Test/compile and compile and is Info by default
    compile / logLevel := Level.Debug,

    // only show warnings and errors on the screen for all tasks (the default is Info)
    //  individual tasks can then be more verbose using the previous setting
    logLevel := Level.Debug
  )