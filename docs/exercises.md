---
layout: default
title: Exercises
---

<div class="gopher-card">
  <h1>GoLearn Exercises 🐹</h1>
  <p>Master Go programming through hands-on exercises! Each exercise is designed to teach specific concepts while building your confidence.</p>
</div>

<div class="gopher-card">
  <h2>Exercise Categories</h2>
  <p>GoLearn exercises are organized into two main categories:</p>
  
  <h3>📚 Concepts (01-37)</h3>
  <p>Learn Go's core language features and syntax through focused exercises.</p>
  
  <h3>🚀 Projects (101-109)</h3>
  <p>Build real applications and apply your Go knowledge to practical projects.</p>
</div>

<div class="gopher-card">
  <h2>Concepts Exercises</h2>
  <p>Start here to learn Go fundamentals:</p>
  
  <div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 20px; margin: 20px 0;">
    <div class="gopher-card" style="margin: 0;">
      <h4>01-10: Basics</h4>
      <ul>
        <li>01_hello - Hello, Go!</li>
        <li>02_values - Values and Types</li>
        <li>03_variables - Variables</li>
        <li>04_constants - Constants</li>
        <li>05_for - For Loops</li>
        <li>06_if_else - If/Else Statements</li>
        <li>07_switch - Switch Statements</li>
        <li>08_arrays - Arrays</li>
        <li>09_slices - Slices</li>
        <li>10_maps - Maps</li>
      </ul>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>11-20: Functions & Structs</h4>
      <ul>
        <li>11_functions - Functions</li>
        <li>12_multi_return - Multiple Return Values</li>
        <li>13_variadic - Variadic Functions</li>
        <li>14_closures - Closures</li>
        <li>15_recursion - Recursion</li>
        <li>16_range_built_in - Range over Built-ins</li>
        <li>17_pointers - Pointers</li>
        <li>18_strings_runes - Strings and Runes</li>
        <li>19_structs - Structs</li>
        <li>20_methods - Methods</li>
      </ul>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>21-30: Advanced Concepts</h4>
      <ul>
        <li>21_interfaces - Interfaces</li>
        <li>22_enums - Enums</li>
        <li>23_struct_embedding - Struct Embedding</li>
        <li>24_generics - Generics</li>
        <li>25_range_iterators - Range over Iterators</li>
        <li>26_errors - Error Handling</li>
        <li>27_custom_errors - Custom Errors</li>
        <li>36_json - JSON Processing</li>
        <li>37_xml - XML Processing</li>
      </ul>
    </div>
  </div>
</div>

<div class="gopher-card">
  <h2>Project Exercises</h2>
  <p>Apply your knowledge by building real applications:</p>
  
  <div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 20px; margin: 20px 0;">
    <div class="gopher-card" style="margin: 0;">
      <h4>101_text_analyzer</h4>
      <p><strong>Difficulty:</strong> Easy</p>
      <p>Build a text analysis tool that counts characters, words, and unique words.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>102_shape_calculator</h4>
      <p><strong>Difficulty:</strong> Medium</p>
      <p>Create a shape area calculator using interfaces and methods.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>103_task_scheduler</h4>
      <p><strong>Difficulty:</strong> Hard</p>
      <p>Build a task scheduler with custom iterators and error handling.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>104_http_server</h4>
      <p><strong>Difficulty:</strong> Easy</p>
      <p>Implement a basic HTTP server that responds to GET requests.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>105_cli_todo_list</h4>
      <p><strong>Difficulty:</strong> Medium</p>
      <p>Create a command-line todo list manager.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>106_simple_chat_app</h4>
      <p><strong>Difficulty:</strong> Medium</p>
      <p>Build a basic client-server chat application.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>107_image_processing_utility</h4>
      <p><strong>Difficulty:</strong> Hard</p>
      <p>Develop an image processing command-line tool.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>108_basic_key_value_store</h4>
      <p><strong>Difficulty:</strong> Hard</p>
      <p>Implement an in-memory key-value store with CRUD operations.</p>
    </div>
    
    <div class="gopher-card" style="margin: 0;">
      <h4>109_epoch</h4>
      <p><strong>Difficulty:</strong> Beginner</p>
      <p>Work with Unix timestamps and time conversion.</p>
    </div>
  </div>
</div>

<div class="gopher-card">
  <h2>How to Use Exercises</h2>
  <p>Here's how to work through GoLearn exercises effectively:</p>
  
  <h3>1. Start with Concepts</h3>
  <p>Begin with the concept exercises (01-37) to learn Go fundamentals. These build upon each other, so it's best to do them in order.</p>
  
  <h3>2. Read the Exercise</h3>
  <p>Each exercise includes:</p>
  <ul>
    <li><strong>Description:</strong> What you need to implement</li>
    <li><strong>Hints:</strong> Helpful guidance to get you started</li>
    <li><strong>Tests:</strong> Automated tests that verify your solution</li>
  </ul>
  
  <h3>3. Work Through the Problem</h3>
  <p>Edit the template file and implement the required functionality. Don't worry if you get stuck - that's part of learning!</p>
  
  <h3>4. Verify Your Solution</h3>
  <p>Use <code>./bin/golearn verify [exercise]</code> to check if your solution is correct.</p>
  
  <h3>5. Move to Projects</h3>
  <p>Once you're comfortable with concepts, try the project exercises to build real applications.</p>
</div>

<div class="gopher-card">
  <h2>Exercise Commands</h2>
  <p>Here are the commands you'll use most often:</p>
  
  <h3>Basic Commands</h3>
  <pre><code># List all exercises
./bin/golearn list

# Verify a specific exercise
./bin/golearn verify 01_hello

# Verify all exercises
./bin/golearn verify

# Get hints for an exercise
./bin/golearn hint 01_hello</code></pre>
  
  <h3>Progress Tracking</h3>
  <pre><code># Show your progress
./bin/golearn progress

# Watch for changes and auto-verify
./bin/golearn watch</code></pre>
  
  <h3>When You're Stuck</h3>
  <pre><code># View the solution (use sparingly!)
./bin/golearn solution 01_hello

# Reset an exercise to start over
./bin/golearn reset 01_hello</code></pre>
</div>

<div class="gopher-card">
  <h2>Learning Path Recommendations</h2>
  <p>Here are some suggested learning paths based on your goals:</p>
  
  <h3>🐹 Complete Beginner</h3>
  <p>Start with exercises 01-10, then move to 11-20, and finally 21-30. Take your time with each concept.</p>
  
  <h3>🚀 Web Developer</h3>
  <p>Focus on 01-20, then jump to 104_http_server, 105_cli_todo_list, and 106_simple_chat_app.</p>
  
  <h3>☁️ DevOps Engineer</h3>
  <p>Complete 01-30, then work on 103_task_scheduler, 108_basic_key_value_store, and 109_epoch.</p>
  
  <h3>🎯 Data Processing</h3>
  <p>Master 01-25, then focus on 101_text_analyzer, 107_image_processing_utility, and 36_json.</p>
</div>

<div class="gopher-card">
  <h2>Need Help?</h2>
  <p>Don't worry if you get stuck! Here are some resources:</p>
  
  <ul>
    <li><strong>Hints:</strong> Use <code>./bin/golearn hint [exercise]</code> for guidance</li>
    <li><strong>Community:</strong> Join the Go community and ask questions</li>
    <li><strong>Documentation:</strong> Check the <a href="https://golang.org/doc/">official Go documentation</a></li>
    <li><strong>Solutions:</strong> Use <code>./bin/golearn solution [exercise]</code> as a last resort</li>
  </ul>
  
  <div style="text-align: center; margin: 30px 0;">
    <a href="/getting-started" class="gopher-btn gopher-btn-primary">Start Learning</a>
    <a href="/contributing" class="gopher-btn gopher-btn-secondary">Help Others Learn</a>
  </div>
</div>
