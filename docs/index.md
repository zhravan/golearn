---
layout: default
title: Home
---

<div class="gopher-card">
  <h1>GoLearn</h1>
  <p>Learn Go through interactive exercises. A Rustlings-style tutorial for the Go programming language.</p>
  
  <div style="text-align: center; margin: 30px 0;">
    <a href="{{ '/getting-started' | relative_url }}" class="gopher-btn gopher-btn-primary">Get Started</a>
    <a href="{{ '/exercises' | relative_url }}" class="gopher-btn gopher-btn-secondary">Browse Exercises</a>
  </div>
</div>

<div class="gopher-card">
  <h2>What is GoLearn?</h2>
  <p>GoLearn provides hands-on exercises that guide you through Go's core concepts, from basic syntax to advanced features.</p>
  
  <ul>
    <li>Interactive coding exercises</li>
    <li>Progressive difficulty levels</li>
    <li>Instant feedback on solutions</li>
    <li>Real project exercises</li>
  </ul>
</div>

<div class="gopher-card">
  <h2>Quick Start</h2>
  <ol>
    <li>Install Go 1.22+</li>
    <li>Install the CLI</li>
    <li>Run your first exercise</li>
  </ol>
  
  <pre><code>go install github.com/zhravan/golearn/cmd/golearn@latest
golearn verify 01_hello</code></pre>
</div>
