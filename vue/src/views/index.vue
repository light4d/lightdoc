<template>
  <div class="z-container">
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
    <div id="sidebar" :style="{
        position: 'fixed'
      }">
      <v-aside :mobileV="mobileV" @handleMobile="handleMobile"></v-aside>
      <transition name="el-zoom-in-center">
        <div id="split-bar" v-if="this.$store.state.splitbar" @mousedown="resize" v-show="mobileV"></div>
      </transition>
    </div>
    <v-markdown
      id="main"
      :style="{
        marginLeft: this.$store.state.splitbar ? '250px' : '20px'
      }"
      :mobileV="mobileV"
      @handleMobile="handleMobile"
    ></v-markdown>
  </div>
</template>
<script>
import vAside from "../components/aside.vue";
import vMarkdown from "../components/markdown.vue";
import smShow from "../components/smShow.vue";
import $ from "jquery";

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
      value: "",
      mobileV: false
    };
  },
  components: {
    vAside,
    vMarkdown,
    smShow
    //BackToTop
  },
  methods: {
    handleMobile(value) {
      if (value !== undefined) {
        console.log("value: " + value);
        this.mobileV = value;
      } else {
        this.mobileV = !this.mobileV;
      }
      console.log("mobileV: " + this.mobileV);
      if (this.mobileV === true)
        document.querySelector("#sidebar").style.backgroundColor = "white";
      else
        document.querySelector("#sidebar").style.backgroundColor =
          "transparent";
    },
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
    },
    resize(e) {
      e.preventDefault();
      var min = 200;
      var max = 800;
      var mainmin = 200;
      $(document).mousemove(function(e) {
        e.preventDefault();
        var x = e.pageX - $("#sidebar").offset().left;
        if (x > min && x < max && e.pageX < $(window).width() - mainmin) {
          $("#sidebar").css("width", x);
          $(".aside-container").css("width", x - 5);
          $("#main").css("margin-left", x);
        }
      });
    }
  },
  mounted() {
    $(document).mouseup(function(e) {
      $(document).unbind("mousemove");
    });
  }
};
</script>
<style lang="stylus" scoped>
.z-container {
  width: 100%;
  position: relative;
  overflow: auto;
  height: 100%;
}

.styleSel {
  position: fixed;
  right: 10px;
  top: 10px;
}

#sidebar {
  // background-color: IndianRed;
  width: 250px;
  height: 100%;
  float: left;
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  transition: all 0.5s ease;
}

#split-bar {
  background-color: #ccc;
  height: 100%;
  width: 6px;
  cursor: col-resize;
  float: right;
  align-self: flex-end;
  z-index: 1000;
}

#main {
  margin-left: 20px;
  // background-color: BurlyWood;
  height: 100%;
}

@media screen and (max-width: 768px) {
  #sidebar {
    display: flex;
    flex-direction: column;
    width: 80%;
    overflow: auto;
    height: 100%;
  }

  #split-bar {
    position: absolute;
  }
}
</style>
