const esbuild = require("esbuild");
const { sassPlugin } = require("esbuild-sass-plugin")

esbuild
    .build({
        entryPoints: ["frontend/Application.jsx", "frontend/style.css"],
        outdir: "public/assets",
        bundle: true,
        minify: true,
        plugins: [sassPlugin()],
    })
    .then(() => console.log("⚡ Build complete! ⚡"))
    .catch(() => process.exit(1));