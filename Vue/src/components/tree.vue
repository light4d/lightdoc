<template>
	<div>
		<ul>
			<li v-for="(item , i) in Floder">{{item.Name}}
				<ul v-if="item.Files">
					<li v-for="(file) in item.Files">{{file}}</li>
				</ul>
			</li>
		</ul>
	</div>
</template>

<script>
export default{
	data(){
		return {
			Floder:[]
		}	
	},
	mounted(){
		this.axios.patch('/api').then((res)=>{
			this.Floder=res.data.Folder
			console.log(res)
		})
	},
}
</script>

<style>
</style>
