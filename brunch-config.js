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
  npm: {
    enabled: true
  },
  modules: {
    autoRequire: {
      "src/maps/public/js/app.js": ["assets/js/app"]
    }
  },
  sourceMaps: false,
  plugins: {
    babel: {presets: ["es2015", "react"]},
    // Do not use ES6 compiler in vendor code
    ignore: [/src\/maps\/\/pubilc\/js\/vendor/],
    autoReload: {
      enabled: true
    },
    sass: {
      mode: "ruby",
      options: {
        sourceMaps: false,
        sourceMapEmbed: false,
        // includePaths: [
        //   "node_modules/material-design-icons"
        // ]
      }
    },
    uglify: {
      mangle: true,
      compress: {
        global_defs: {
          DEBUG: false
        }
      }
    }
  }
};
