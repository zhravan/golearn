---
layout: default
title: Contributing
---

<div class="gopher-card">
  <h1>Contributing to GoLearn 🐹</h1>
  <p>GoLearn is an open-source project built by the community, for the community. We welcome contributions of all kinds!</p>
</div>

<div class="gopher-card">
  <h2>Ways to Contribute</h2>
  <p>There are many ways you can help make GoLearn better:</p>
  
  <h3>🐛 Bug Reports</h3>
  <p>Found a bug? Help us fix it by reporting it on GitHub Issues. Include:</p>
  <ul>
    <li>Clear description of the problem</li>
    <li>Steps to reproduce the issue</li>
    <li>Expected vs actual behavior</li>
    <li>Your system information (OS, Go version, etc.)</li>
  </ul>
  
  <h3>💡 Feature Requests</h3>
  <p>Have an idea for a new feature or exercise? We'd love to hear it! Please include:</p>
  <ul>
    <li>Clear description of the feature</li>
    <li>Why it would be useful</li>
    <li>How it fits with GoLearn's goals</li>
  </ul>
  
  <h3>📝 Documentation</h3>
  <p>Help improve our documentation by:</p>
  <ul>
    <li>Fixing typos and grammar</li>
    <li>Adding examples and explanations</li>
    <li>Improving clarity and structure</li>
    <li>Translating to other languages</li>
  </ul>
  
  <h3>🎯 New Exercises</h3>
  <p>Create new exercises to help others learn Go:</p>
  <ul>
    <li>Concept exercises for specific Go features</li>
    <li>Project exercises for real-world applications</li>
    <li>Exercises for different difficulty levels</li>
  </ul>
</div>

<div class="gopher-card">
  <h2>Getting Started</h2>
  <p>Ready to contribute? Here's how to get started:</p>
  
  <h3>1. Fork the Repository</h3>
  <p>Click the "Fork" button on GitHub to create your own copy of GoLearn.</p>
  
  <h3>2. Clone Your Fork</h3>
  <pre><code>git clone https://github.com/YOUR_USERNAME/golearn.git
cd golearn</code></pre>
  
  <h3>3. Create a Branch</h3>
  <pre><code>git checkout -b your-feature-name</code></pre>
  
  <h3>4. Make Your Changes</h3>
  <p>Make your changes and test them thoroughly.</p>
  
  <h3>5. Test Your Changes</h3>
  <pre><code># Run the full test suite
go test ./...

# Test specific exercises
./bin/golearn verify

# Test with solutions
./bin/golearn verify 01_hello --solution</code></pre>
  
  <h3>6. Commit and Push</h3>
  <pre><code>git add .
git commit -m "Add your descriptive commit message"
git push origin your-feature-name</code></pre>
  
  <h3>7. Create a Pull Request</h3>
  <p>Open a pull request on GitHub with a clear description of your changes.</p>
</div>

<div class="gopher-card">
  <h2>Creating New Exercises</h2>
  <p>Want to add a new exercise? Here's how:</p>
  
  <h3>1. Choose a Topic</h3>
  <p>Pick a Go concept or project that would be valuable for learners. Make sure it's not already covered!</p>
  
  <h3>2. Create the Exercise Structure</h3>
  <p>Each exercise needs:</p>
  <ul>
    <li><strong>Template file:</strong> Incomplete code for students to fill in</li>
    <li><strong>Test file:</strong> Tests that verify the solution</li>
    <li><strong>Solution file:</strong> Complete working implementation</li>
    <li><strong>Catalog entry:</strong> Metadata in catalog.yaml</li>
  </ul>
  
  <h3>3. Write the Template</h3>
  <p>Create a template file with:</p>
  <ul>
    <li>Clear comments explaining what needs to be implemented</li>
    <li>Function signatures and basic structure</li>
    <li>Helpful variable names and type hints</li>
  </ul>
  
  <h3>4. Write the Tests</h3>
  <p>Create comprehensive tests that:</p>
  <ul>
    <li>Cover all the required functionality</li>
    <li>Test edge cases and error conditions</li>
    <li>Provide clear error messages when tests fail</li>
  </ul>
  
  <h3>5. Write the Solution</h3>
  <p>Create a complete, working solution that passes all tests.</p>
  
  <h3>6. Add to Catalog</h3>
  <p>Update <code>internal/exercises/catalog.yaml</code> with your exercise metadata.</p>
</div>

<div class="gopher-card">
  <h2>Exercise Guidelines</h2>
  <p>When creating exercises, follow these guidelines:</p>
  
  <h3>🐹 Keep It Fun</h3>
  <p>Make exercises engaging and enjoyable. Use creative examples and interesting problems.</p>
  
  <h3>📚 Educational Value</h3>
  <p>Each exercise should teach something specific about Go. Avoid exercises that are just busy work.</p>
  
  <h3>🎯 Appropriate Difficulty</h3>
  <p>Make sure the difficulty level matches the target audience. Provide hints for complex exercises.</p>
  
  <h3>✅ Test Everything</h3>
  <p>Write comprehensive tests that cover all functionality and edge cases.</p>
  
  <h3>📖 Clear Documentation</h3>
  <p>Provide clear descriptions, hints, and examples to help learners understand what to do.</p>
</div>

<div class="gopher-card">
  <h2>Code Style</h2>
  <p>Follow these coding standards:</p>
  
  <h3>Go Code Style</h3>
  <ul>
    <li>Follow <a href="https://golang.org/doc/effective_go.html">Effective Go</a> guidelines</li>
    <li>Use <code>gofmt</code> to format your code</li>
    <li>Run <code>go vet</code> to check for common mistakes</li>
    <li>Use meaningful variable and function names</li>
  </ul>
  
  <h3>Documentation Style</h3>
  <ul>
    <li>Use clear, concise language</li>
    <li>Provide examples where helpful</li>
    <li>Keep explanations beginner-friendly</li>
    <li>Use consistent formatting</li>
  </ul>
</div>

<div class="gopher-card">
  <h2>Review Process</h2>
  <p>All contributions go through a review process:</p>
  
  <h3>1. Automated Checks</h3>
  <p>Your pull request will be automatically tested for:</p>
  <ul>
    <li>Code formatting and style</li>
    <li>Test coverage and correctness</li>
    <li>Build and compilation</li>
  </ul>
  
  <h3>2. Community Review</h3>
  <p>Community members will review your changes for:</p>
  <ul>
    <li>Code quality and correctness</li>
    <li>Educational value and clarity</li>
    <li>Consistency with project goals</li>
  </ul>
  
  <h3>3. Feedback and Iteration</h3>
  <p>Be prepared to make changes based on feedback. This is normal and helps improve the project!</p>
</div>

<div class="gopher-card">
  <h2>Community Guidelines</h2>
  <p>Help us maintain a welcoming and inclusive community:</p>
  
  <h3>🤝 Be Respectful</h3>
  <p>Treat everyone with respect and kindness. We're all here to learn and help others.</p>
  
  <h3>💬 Be Constructive</h3>
  <p>Provide helpful feedback and suggestions. Focus on improving the code, not criticizing the person.</p>
  
  <h3>🎓 Be Patient</h3>
  <p>Remember that everyone is at different skill levels. Be patient with questions and mistakes.</p>
  
  <h3>🌟 Be Encouraging</h3>
  <p>Celebrate successes and encourage others in their learning journey.</p>
</div>

<div class="gopher-card">
  <h2>Need Help?</h2>
  <p>Don't hesitate to ask for help:</p>
  
  <ul>
    <li><strong>GitHub Issues:</strong> Ask questions about contributing</li>
    <li><strong>Discussions:</strong> Join community discussions</li>
    <li><strong>Code Review:</strong> Ask for help with your pull request</li>
    <li><strong>Documentation:</strong> Check existing docs and examples</li>
  </ul>
  
  <div style="text-align: center; margin: 30px 0;">
    <a href="https://github.com/your-username/golearn/issues" class="gopher-btn gopher-btn-primary">Report an Issue</a>
    <a href="https://github.com/your-username/golearn/discussions" class="gopher-btn gopher-btn-secondary">Join Discussion</a>
  </div>
</div>

<div class="gopher-card">
  <h2>Thank You!</h2>
  <p>Thank you for contributing to GoLearn! Every contribution, no matter how small, helps make Go learning more accessible and fun for everyone.</p>
  
  <div style="text-align: center; margin: 30px 0;">
    <div style="font-size: 3rem; margin: 20px 0;">🐹</div>
    <p>Happy coding, and welcome to the GoLearn community!</p>
  </div>
</div>
