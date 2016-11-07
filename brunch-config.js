// See http://brunch.io for documentation.

module.exports = {
  paths: {
    public: "src/maps/public"
  },
  files: {
    javascripts: {
      joinTo: {
        "assets/js/app.js": "js/app.js"
      }
    },
    stylesheets: {joinTo: "css/app.css"}
    // templates: {joinTo: "app.js"}
  },
  plugins: {
    babel: {presets: ["es2015"]}
  }
};
