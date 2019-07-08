<template>
  <div id="app" class="app">
    <router-view />
  </div>
</template>

<script>
import { Backtop } from "element-ui";
export default {
  name: "App"
};
</script>

<style lang="stylus" scoped>
#app {
  height: 100% !important;
}
</style>
