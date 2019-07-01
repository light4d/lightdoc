<template>
  <div class="z-container">
    <div class="content">
      <div class="styleSel">
        <el-select
          v-model="value"
          size="mini"
          placeholder="样式选择"
          style="width: 100px"
          @change="getValue"
        >
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

export default {
  data() {
    return {
      options: [
        {
          value: "1",
          label: "Green",
          rgb: "#009688"
        },
        {
          value: "2",
          label: "Blue",
          rgb: "#409EFF"
        },
        {
          value: "3",
          label: "Black",
          rgb: "#303133"
        }
      ],
      value: ""
    };
  },
  components: {
    vAside,
    vMarkdown,
    smShow
  },
  methods: {
    getValue(color) {
      var elements = [];
      elements[0] = document.querySelectorAll("h1");
      elements[1] = document.querySelectorAll("h2");
      elements[2] = document.querySelectorAll("h3");
      elements[3] = document.querySelectorAll("h4");
      console.log(elements);
      for (var i = 0; i < elements.length; i++) {
        for (var n = 0; n < elements[i].length; n++) {
          elements[i][n].style.color = color.rgb;
        }
      }
    }
  }
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
  right: 10px;
  top: 10px;
}

.btn-to-top {
  width: 40px;
  height: 40px;
  padding: 10px 10px;
  border-radius: 50%;
  color: '#666';
  font-size: 20px;
  line-height: 20px;
  transition: all 0.3s ease-in-out;
}

.vue-back-to-top {
  background-color: '#666' !important;
}
</style>
