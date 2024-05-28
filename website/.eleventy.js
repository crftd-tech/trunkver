const prismCss = require.resolve("prismjs/themes/prism.css");

const syntaxHighlight = require("@11ty/eleventy-plugin-syntaxhighlight");

module.exports = function (eleventyConfig) {
  eleventyConfig.addPlugin(syntaxHighlight);
  eleventyConfig.addPassthroughCopy({
    [prismCss]: "prism.css",
	"style.css": "style.css",
  });
};
