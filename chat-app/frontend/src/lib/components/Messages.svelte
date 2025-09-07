<script>
	const prop = $props();
	let data = $state();
	if (prop.msg === '' || prop.msg != null) {
		data = JSON.parse(prop.msg);
        
	}
</script>

<div class="p-3" class:text-right={data.name === prop.name.data.username} class:text-left={data.name !== prop.name.data.username}>
    <span class="text-sm text-gray-500">~{data.name}</span>
    <p class="mb-4">{data.message}</p>
</div>
