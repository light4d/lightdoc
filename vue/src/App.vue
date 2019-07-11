<template>
  <div id="app" class="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
  mounted() {
    document.title = "Hello Lightdoc";
  }
};
</script>

<style lang="stylus" scoped>
#app {
  height: 100% !important;
}
</style>
