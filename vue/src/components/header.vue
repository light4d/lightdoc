<template>
	<div class="logo-container">
		<div class="logo">
			<img v-if="url" src='http://localhost:8000/logo.jpg' />
		</div>
		<div class="fixed" @click="handlefixed">
			<img src="@/assets/iconfont/guding.svg" alt="" :class="{'animated':fixed==true}" />
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
			urllogo:'http://localhost/logo.jpg',
			fixed: false
		}
	},
	mounted(){
		// this.axios.options('http://localhost:8000').then((res)=>{
		// 	let arr = Object.getOwnPropertyNames(res.data); //arr=["one", "two"]
			
		// 	let Url = arr.map(function(i){return res.data[i]})
		// 	this.url = 'http://localhost:8000'+Url[0]
			
		// })
	},
	methods: {
		handlefixed(){
			this.fixed = !this.fixed
			this.$emit('fixed' , this.fixed)
		}
	},
}
</script>

<style lang="stylus" scoped>
	.logo-container
		height:50px
		width:100%
		background:#fff
		margin-left:35%
		display : flex
		.logo
			width:84px
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
		.fixed
			width : 100%
			margin-left : 34px
			height : 50px
			line-height : 50px
			cursor : pointer
		img
			width : 20px
			height : 20px
		.animated
			transform : rotate(-45deg)
</style>
