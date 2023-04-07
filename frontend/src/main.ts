// Your selected Skeleton theme:
import '@skeletonlabs/skeleton/themes/theme-skeleton.css';

// This contains the bulk of Skeletons required styles:
import '@skeletonlabs/skeleton/styles/all.css';

import "./app.postcss";
import "./app.scss";


import App from "./App.svelte";

const app = new App({
  target: document.getElementById("app"),
});

export default app;
