<template>
	<div class="container">
		<div class="logo">
			<img v-if="url" :src=url />
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
			url:''
		}
	},
	mounted(){
		this.axios.options('/api').then((res)=>{
			let arr = Object.getOwnPropertyNames(res.data); //arr=["one", "two"]
			
			let Url = arr.map(function(i){return res.data[i]})
			this.url = '/api'+Url[0]
			
		})
	}
}
</script>

<style lang="stylus" scoped>
	.container
		height:50px
		width:100%
		background:#fff
		display:flex
		flex-direction:row
		justify-content:center
		align-items:center
		.logo
			width:50px
			height:50px
			img
				width:100%
				height:100%
		@media screen and (min-width: 768px)
			.hidden
				display:none
</style>
