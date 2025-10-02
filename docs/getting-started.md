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
  <pre><code>go install github.com/zhravan/golearn/cmd/golearn@latest
golearn --help</code></pre>
</div>

<div class="gopher-card">
  <h2>First Exercise</h2>
  <ol>
    <li>List exercises: <code>golearn list</code></li>
    <li>Start first exercise: <code>golearn verify 01_hello</code></li>
    <li>Get hints: <code>golearn hint 01_hello</code></li>
    <li>Edit the template file and implement the solution</li>
    <li>Verify your solution: <code>golearn verify 01_hello</code></li>
  </ol>
</div>

<div class="gopher-card">
  <h2>Getting Started as a Learner (CLI only)</h2>
  <ol>
    <li>List available exercises: <code>golearn list</code></li>
    <li>Initialize local templates (optional): <code>golearn init</code></li>
    <li>Try an exercise (expect failure until you implement it): <code>golearn verify 01_hello</code></li>
    <li>Read a hint: <code>golearn hint 01_hello</code></li>
    <li>Edit the template file for the exercise, then re-run <code>verify</code></li>
    <li>See your overall progress: <code>golearn progress</code></li>
    <li>Optionally view the reference solution: <code>golearn solution 01_hello</code></li>
    <li>Reset an exercise to its starter state: <code>golearn reset 01_hello</code></li>
  </ol>
</div>

<div class="gopher-card">
  <h2>Getting Started as a Contributor</h2>
  <ol>
    <li>Read the contribution guide: <a href="{{ '/contributing' | relative_url }}">Contributing</a></li>
    <li>Fork and clone the repo, create a feature branch</li>
    <li>Build the CLI locally: <code>go build -o bin/golearn ./cmd/golearn && chmod +x bin/golearn</code></li>
    <li>Develop your exercise/update under <code>internal/exercises</code></li>
    <li>Verify templates (incomplete templates should fail): <code>./bin/golearn verify</code></li>
    <li>Verify specific exercise with solution (should pass): <code>./bin/golearn verify 01_hello --solution</code></li>
    <li>Run auxiliary checks locally: <code>go fmt ./...</code> and <code>go vet ./...</code></li>
    <li>Open a Pull Request</li>
  </ol>
</div>

<div class="gopher-card">
  <h2>Commands</h2>
  <ul>
    <li><code>golearn list</code> - List all exercises</li>
    <li><code>golearn verify [exercise]</code> - Verify an exercise</li>
    <li><code>golearn hint [exercise]</code> - Get hints</li>
    <li><code>golearn progress</code> - Show progress</li>
    <li><code>golearn solution [exercise]</code> - View solution</li>
  </ul>

  <h3>CLI Commands</h3>
  <table>
    <thead>
      <tr>
        <th>Command</th>
        <th>Description</th>
        <th>Example</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td><code>golearn --help</code></td>
        <td>Show help and available commands</td>
        <td><code>./bin/golearn --help</code></td>
      </tr>
      <tr>
        <td><code>golearn list</code></td>
        <td>List all available exercises</td>
        <td><code>./bin/golearn list</code></td>
      </tr>
      <tr>
        <td><code>./bin/golearn verify [exercise]</code></td>
        <td>Verify all templates or a specific exercise</td>
        <td>
          <div><code>golearn verify</code></div>
          <div><code>golearn verify 01_hello</code></div>
        </td>
      </tr>
      <tr>
        <td><code>./bin/golearn verify [exercise] --solution</code></td>
        <td>Verify using the reference solution (should pass)</td>
        <td><code>golearn verify 01_hello --solution</code></td>
      </tr>
      <tr>
        <td><code>./bin/golearn hint [exercise]</code></td>
        <td>Show hints for a specific exercise</td>
        <td><code>golearn hint 01_hello</code></td>
      </tr>
      <tr>
        <td><code>./bin/golearn solution [exercise]</code></td>
        <td>View the reference solution (may prompt for confirmation)</td>
        <td><code>golearn solution 01_hello</code></td>
      </tr>
      <tr>
        <td><code>golearn progress</code></td>
        <td>Display your overall progress</td>
        <td><code>./bin/golearn progress</code></td>
      </tr>
      <tr>
        <td><code>./bin/golearn reset [exercise]</code></td>
        <td>Reset an exercise to its starter template</td>
        <td><code>golearn reset 01_hello</code></td>
      </tr>
      <tr>
        <td><code>golearn init</code></td>
        <td>Initialize local exercise templates</td>
        <td><code>./bin/golearn init</code></td>
      </tr>
      <tr>
        <td><code>./bin/golearn publish [--dry-run]</code></td>
        <td>Run publish routine (use <code>--dry-run</code> to preview)</td>
        <td><code>golearn publish --dry-run</code></td>
      </tr>
      <tr>
        <td><code>golearn watch</code></td>
        <td>Watch files and re-verify on changes</td>
        <td><code>./bin/golearn watch</code></td>
      </tr>
    </tbody>
  </table>

  <h3>Global Flags</h3>
  <table>
    <thead>
      <tr>
        <th>Flag</th>
        <th>Description</th>
        <th>Example</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td><code>--no-color</code></td>
        <td>Disable colored output</td>
        <td><code>golearn --no-color list</code></td>
      </tr>
      <tr>
        <td><code>--theme=&lt;name&gt;</code></td>
        <td>Set output theme (<code>high-contrast</code>, <code>monochrome</code>)</td>
        <td><code>golearn --theme=high-contrast list</code></td>
      </tr>
    </tbody>
  </table>
</div>
