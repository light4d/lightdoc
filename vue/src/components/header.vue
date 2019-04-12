<template>
	<div class="logo-container">
		<div class="logo">
			<img v-if="url" :src='urllogo?urllogo:url' />
		</div>
	</div>
</template>

<script>
import 'element-ui/lib/theme-chalk/display.css';
export default {
	name:'vHeader',
	components: {
	},
	data () {
		return {
			url:'@/assets/logo.jpg',
			urllogo:'http://localhost/logo.jpg'
		}
	},
	mounted(){
		// this.axios.options('http://localhost:8000').then((res)=>{
		// 	let arr = Object.getOwnPropertyNames(res.data); //arr=["one", "two"]
			
		// 	let Url = arr.map(function(i){return res.data[i]})
		// 	this.url = 'http://localhost:8000'+Url[0]
			
		// })
	}
}
</script>

<style lang="stylus" scoped>
	.logo-container
		height:50px
		width:100%
		background:#fff
		margin-left:35%
		.logo
			width:50px
			height:50px
			border:1px solid #fff
			border-radius:50%
			overflow: hidden
			img
				width:100%
				height:100%
		@media screen and (min-width: 768px)
			.hidden
				display:none
</style>
