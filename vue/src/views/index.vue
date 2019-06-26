<template>
  <div class="z-container">
    <vm-back-top></vm-back-top>
    <div class="content">
      <div class="styleSel">
        <el-select v-model="value" placeholder="样式选择" @change="getValue">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item"></el-option>
        </el-select>
      </div>
      <v-aside></v-aside>
      <v-markdown></v-markdown>
    </div>
  </div>
</template>
<script>
import vAside from "../components/aside.vue";
import vMarkdown from "../components/markdown.vue";
import smShow from "../components/smShow.vue";
import Vue from "vue";
import VmBackTop from "vue-multiple-back-top";

import "element-ui/lib/theme-chalk/display.css";
export default {
  data() {
    return {
      options: [
        {
          value: "1",
          label: "Green"
        },
        {
          value: "2",
          label: "Blue"
        },
        {
          value: "3",
          label: "Black"
        }
      ],
      value: ""
    };
  },
  components: {
    vAside,
    vMarkdown,
    smShow,
    VmBackTop
  },
  methods: {}
};
</script>
<style lang="stylus" scoped>
.z-container {
  width: 100%;
  position: relative;
  display: flex;
  overflow: auto;
  flex-direction: column;
  height: 100%;

  .content {
    display: flex;
    width: 100%;
    flex-direction: row;
    height: 100%;
    overflow: hidden;
  }

  @media screen and (max-width: 768px) {
    .smShow {
      background: #CCC;
      width: 100%;
      height: auto;
      position: absolute;
      top: 7%;
      left: 0;
      display: block;
    }
  }
}

.styleSel {
  float: right;
  position: absolute;
  right: 100px;
  top: 20px;
}
</style>
