<script>
    const prop = $props();    
    let data = $state();
    if (prop.msg === "" || prop.msg != null) {
        data = JSON.parse(prop.msg)        
    }
    
</script>

<div class="text-{prop.direction === "right" ? "right" : "left"}">
    <span class="text-sm text-gray-500">~{data.name}</span>
    <p class="mb-4">{data.message}</p>
</div>