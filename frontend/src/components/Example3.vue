<template>
	<div id="example-3">
		<button v-on:click="say('Vue')">Say Vue</button>
		<button v-on:click="say('JS')">Say JS</button>
	</div>
</template>

<script>
export default{
	el: "#example-3",
	methods: {
		say: function(message) {
			alert(message)
		}
	}
}
</script>