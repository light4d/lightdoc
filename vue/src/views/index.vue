<template>
	<div class="container">
	    <!-- <v-header></v-header> -->
		<div class="content">
			<v-aside  ></v-aside>
			<v-markdown ></v-markdown>
		</div>
		
	</div>
</template>
<script>
import vHeader from '../components/header.vue'
import vAside from '../components/aside.vue'
import vMarkdown from '../components/markdown.vue'
import smShow from '../components/smShow.vue'
import 'element-ui/lib/theme-chalk/display.css';
export default {
	data (){
		return {}
	},
  components: {
	vAside,
	vMarkdown,
	smShow
  },
  methods: {
		
	}
}
</script>
<style lang="stylus" scoped>
	.container
		position:relative
		display:flex
		overflow :auto
		flex-direction:column
		height:100%
		.content
			display:flex
			flex-direction:row
			height:100%
			overflow :hidden
		@media screen and (max-width:768px)
			.smShow
				background:#CCC
				width:100%
				height:auto
				position:absolute
				top:7%
				left:0
				display:block
</style>
