# GoLearn Documentation

This directory contains the Jekyll-based documentation website for GoLearn.

## Local Development

To run the documentation site locally:

1. Install Ruby and Bundler
2. Install dependencies:

   ```bash
   cd docs
   bundle install
   ```

3. Start the Jekyll server:

   ```bash
   bundle exec jekyll serve
   ```

4. Open <http://localhost:4000> in your browser

## Deployment

The site is automatically deployed to GitHub Pages when changes are pushed to the main branch.

## Structure

- `_config.yml` - Jekyll configuration
- `_layouts/` - HTML templates
- `_includes/` - Reusable components
- `assets/` - CSS, JS, and other assets
- `*.md` - Markdown pages

## Theme

The site uses a custom gopher-themed design with:

- Playful gopher animations
- Go-inspired color scheme
- Responsive design
- Clean, readable typography
