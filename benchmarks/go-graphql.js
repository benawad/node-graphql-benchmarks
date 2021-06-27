"use strict";
const { exec } = require("child_process");
const path = require("path");

const forked = exec("/usr/local/go/bin/go run server.go", { cwd: path.join(__dirname, "..", "other-benchmarks/go-gql/") }, (error, stdout, stderr) => {
  if (error) {
    console.log(`error: ${error.message}`);
    return;
  }
  if (stderr) {
    console.log(`stderr: ${stderr}`);
    return;
  }
  console.log(`stdout: ${stdout}`);
});
console.log(path.join(__dirname, "..", "other-benchmarks/go-gql/server.go"))
forked.on("exit", () => console.log("exited"))