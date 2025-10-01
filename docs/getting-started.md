---
layout: default
title: Getting Started
---

<div class="gopher-card">
  <h1>Getting Started with GoLearn 🐹</h1>
  <p>Welcome to your Go learning journey! This guide will help you set up GoLearn and start your first exercises.</p>
</div>

<div class="gopher-card">
  <h2>Prerequisites</h2>
  <p>Before you begin, make sure you have the following installed:</p>
  
  <h3>1. Go Programming Language</h3>
  <p>GoLearn requires Go 1.22 or later. If you don't have Go installed:</p>
  
  <h4>Install Go:</h4>
  <ul>
    <li><strong>macOS:</strong> <code>brew install go</code> or download from <a href="https://golang.org/dl/">golang.org</a></li>
    <li><strong>Linux:</strong> <code>sudo apt install golang-go</code> or download from <a href="https://golang.org/dl/">golang.org</a></li>
    <li><strong>Windows:</strong> Download and install from <a href="https://golang.org/dl/">golang.org</a></li>
  </ul>
  
  <h4>Verify Installation:</h4>
  <pre><code>go version</code></pre>
  <p>You should see something like: <code>go version go1.22.0 linux/amd64</code></p>
</div>

<div class="gopher-card">
  <h2>Installation</h2>
  <p>Now let's get GoLearn up and running:</p>
  
  <h3>1. Clone the Repository</h3>
  <pre><code>git clone https://github.com/your-username/golearn.git
cd golearn</code></pre>
  
  <h3>2. Build GoLearn</h3>
  <pre><code>go build -o bin/golearn ./cmd/golearn
chmod +x bin/golearn</code></pre>
  
  <h3>3. Verify Installation</h3>
  <pre><code>./bin/golearn --help</code></pre>
  <p>You should see the GoLearn help message with available commands.</p>
</div>

<div class="gopher-card">
  <h2>Your First Exercise</h2>
  <p>Let's start with the classic "Hello, Go!" exercise:</p>
  
  <h3>1. List Available Exercises</h3>
  <pre><code>./bin/golearn list</code></pre>
  <p>This shows all available exercises organized by concepts and projects.</p>
  
  <h3>2. Start the First Exercise</h3>
  <pre><code>./bin/golearn verify 01_hello</code></pre>
  <p>This will show you the exercise and its current status (it should fail initially).</p>
  
  <h3>3. Get a Hint</h3>
  <pre><code>./bin/golearn hint 01_hello</code></pre>
  <p>Get helpful hints to guide you through the exercise.</p>
  
  <h3>4. Work on the Exercise</h3>
  <p>Open the exercise file and implement the required functionality:</p>
  <pre><code># Edit the template file
nano internal/exercises/templates/01_hello/hello.go</code></pre>
  
  <h3>5. Verify Your Solution</h3>
  <pre><code>./bin/golearn verify 01_hello</code></pre>
  <p>Run this command to check if your solution is correct.</p>
</div>

<div class="gopher-card">
  <h2>GoLearn Commands</h2>
  <p>Here are the main commands you'll use:</p>
  
  <h3>Core Commands</h3>
  <ul>
    <li><code>./bin/golearn list</code> - List all available exercises</li>
    <li><code>./bin/golearn verify [exercise]</code> - Verify an exercise (or all exercises)</li>
    <li><code>./bin/golearn hint [exercise]</code> - Get hints for an exercise</li>
    <li><code>./bin/golearn progress</code> - Show your learning progress</li>
  </ul>
  
  <h3>Advanced Commands</h3>
  <ul>
    <li><code>./bin/golearn solution [exercise]</code> - View the solution (use sparingly!)</li>
    <li><code>./bin/golearn reset [exercise]</code> - Reset an exercise to its original state</li>
    <li><code>./bin/golearn watch</code> - Watch for file changes and auto-verify</li>
    <li><code>./bin/golearn init</code> - Initialize exercises in a new directory</li>
  </ul>
</div>

<div class="gopher-card">
  <h2>Learning Tips</h2>
  <p>Here are some tips to make the most of your GoLearn experience:</p>
  
  <h3>🐹 Take Your Time</h3>
  <p>Don't rush through exercises. Take time to understand each concept before moving on.</p>
  
  <h3>🔍 Read the Tests</h3>
  <p>Look at the test files to understand what your code should do. Tests are great documentation!</p>
  
  <h3>💡 Use Hints Wisely</h3>
  <p>Try to solve exercises on your own first, but don't hesitate to use hints when you're stuck.</p>
  
  <h3>🔄 Practice Regularly</h3>
  <p>Consistent practice is key to mastering Go. Try to do a few exercises each day.</p>
  
  <h3>🤝 Ask for Help</h3>
  <p>Join the Go community, ask questions, and help others. Learning together is more fun!</p>
</div>

<div class="gopher-card">
  <h2>Next Steps</h2>
  <p>Ready to dive deeper? Here's what to do next:</p>
  
  <ol>
    <li><strong>Complete the Basics:</strong> Work through exercises 01-10 to learn Go fundamentals</li>
    <li><strong>Explore Advanced Topics:</strong> Try exercises 11-27 for more complex concepts</li>
    <li><strong>Build Projects:</strong> Tackle the project exercises (101-109) to build real applications</li>
    <li><strong>Contribute:</strong> Help improve GoLearn by contributing exercises or fixing bugs</li>
  </ol>
  
  <div style="text-align: center; margin: 30px 0;">
    <a href="/exercises" class="gopher-btn gopher-btn-primary">Browse All Exercises</a>
    <a href="/contributing" class="gopher-btn gopher-btn-secondary">Contribute to GoLearn</a>
  </div>
</div>
