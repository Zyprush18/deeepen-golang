<script>
    const prop = $props()
    console.log(prop);
    
    let data = $state();
    if (prop.msg != '' || prop.msg != null) {
        data = JSON.parse(prop.msg)        
    }
    
    
</script>
<span class="text-sm text-gray-300">~{data.name}</span>
<p class="mb-4">{data.message}</p>