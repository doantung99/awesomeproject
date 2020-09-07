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
			.get('https://api.coindesk.com/v1/bpi/currentprice.json')
			.then(response => (this.info = response.data.bpi))
	}
}
</script>