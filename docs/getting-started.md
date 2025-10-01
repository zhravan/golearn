---
layout: default
title: Getting Started
---

<div class="gopher-card">
  <h1>Getting Started</h1>
  <p>Set up GoLearn and start your first exercise.</p>
</div>

<div class="gopher-card">
  <h2>Prerequisites</h2>
  <p>Install Go 1.22 or later:</p>
  <ul>
    <li><strong>macOS:</strong> <code>brew install go</code></li>
    <li><strong>Linux:</strong> <code>sudo apt install golang-go</code></li>
    <li><strong>Windows:</strong> Download from <a href="https://golang.org/dl/">golang.org</a></li>
  </ul>
  
  <pre><code>go version</code></pre>
</div>

<div class="gopher-card">
  <h2>Installation</h2>
  <pre><code>git clone https://github.com/your-username/golearn.git
cd golearn
go build -o bin/golearn ./cmd/golearn
chmod +x bin/golearn</code></pre>
</div>

<div class="gopher-card">
  <h2>First Exercise</h2>
  <ol>
    <li>List exercises: <code>./bin/golearn list</code></li>
    <li>Start first exercise: <code>./bin/golearn verify 01_hello</code></li>
    <li>Get hints: <code>./bin/golearn hint 01_hello</code></li>
    <li>Edit the template file and implement the solution</li>
    <li>Verify your solution: <code>./bin/golearn verify 01_hello</code></li>
  </ol>
</div>

<div class="gopher-card">
  <h2>Commands</h2>
  <ul>
    <li><code>./bin/golearn list</code> - List all exercises</li>
    <li><code>./bin/golearn verify [exercise]</code> - Verify an exercise</li>
    <li><code>./bin/golearn hint [exercise]</code> - Get hints</li>
    <li><code>./bin/golearn progress</code> - Show progress</li>
    <li><code>./bin/golearn solution [exercise]</code> - View solution</li>
  </ul>
</div>
