<template>
	<div class="container">
	    <v-header></v-header>
		<div class="content">
			<v-aside class="hidden-xs-only"></v-aside>
			<v-markdown></v-markdown>
		</div>
		<sm-show class="smShow hidden-sm-and-up"></sm-show>
	</div>
</template>
<script>
import vHeader from '../components/header.vue'
import vAside from '../components/aside.vue'
import vMarkdown from '../components/markdown.vue'
import smShow from '../components/smShow.vue'
import 'element-ui/lib/theme-chalk/display.css';
export default {
  components: {
	vHeader,
	vAside,
	vMarkdown,
	smShow
  },
  methods: {}
}
</script>
<style lang="stylus" scoped>
	.container
		position:relative
		display:flex
		flex-direction:column
		height:100%
		.content
			display:flex
			flex-direction:row
			height:100%
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
