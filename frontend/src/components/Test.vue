<template>
	<div id="test">
		{{ info }}
	</div>
</template>

<script>
import axios from 'axios'

export default{
	name: 'Test',
	data(){
		return{
			info: null
		}	
	},
	mounted(){
		axios
			.get('http://localhost:8081/reqlist')
			.then(response => (this.info = response.data))
	}
}
</script>