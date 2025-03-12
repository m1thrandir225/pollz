<script lang="ts">
	let { children } = $props();
</script>

<main>
	<h1>About Layout Wrapper</h1>
	<div>
		{@render children()}
	</div>
</main>
