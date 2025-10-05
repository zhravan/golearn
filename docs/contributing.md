---
layout: default
title: Contributing
---

<div class="gopher-card">
  <h1>Contributing</h1>
  <p>Help improve GoLearn by contributing exercises, fixing bugs, or improving documentation.</p>
</div>

<div class="gopher-card">
  <h2>Ways to Contribute</h2>
  <ul>
    <li><strong>Bug Reports:</strong> Report issues on GitHub</li>
    <li><strong>New Exercises:</strong> Create exercises for Go concepts</li>
    <li><strong>Documentation:</strong> Improve guides and examples</li>
    <li><strong>Code:</strong> Fix bugs and add features</li>
  </ul>
</div>

<div class="gopher-card">
  <h2>Getting Started</h2>
  <ol>
    <li>Fork the repository on GitHub</li>
    <li>Clone your fork: <code>git clone https://github.com/YOUR_USERNAME/golearn.git</code></li>
    <li>Create a branch: <code>git checkout -b your-feature-name</code></li>
    <li>Make your changes and test them</li>
    <li>Commit and push: <code>git commit -m "Your message"</code></li>
    <li>Create a pull request on GitHub</li>
  </ol>
</div>

<div class="gopher-card">
  <h2>Creating Exercises</h2>
  <p>Each exercise needs:</p>
  <ul>
    <li>Template file (incomplete code for learners to complete)</li>
    <li>Test file (verification tests)</li>
    <li>Solution file (complete implementation, base code is already in place which will run the test cases written for template to picked for solutions as well)</li>
    <li>Catalog entry in <code>catalog.yaml</code></li>
  </ul>
</div>

<div class="gopher-card">
  <h2>Code Style</h2>
  <ul>
    <li>Follow <a href="https://golang.org/doc/effective_go.html">Effective Go</a> guidelines</li>
    <li>Use <code>gofmt</code> to format code</li>
    <li>Run <code>go vet</code> to check for issues</li>
    <li>Write clear, descriptive commit messages</li>
  </ul>
</div>

<div class="gopher-card">
  <h2>Testing</h2>
  <p>Before submitting your changes, build locally and verify exercises using the CLI:</p>
  <pre><code># Build the CLI
go build -o bin/golearn ./cmd/golearn
chmod +x bin/golearn

# Verify template exercises (incomplete templates should fail)
./bin/golearn verify

# Verify a specific exercise with its solution (should pass)
./bin/golearn verify 01_hello --solution</code></pre>
</div>
