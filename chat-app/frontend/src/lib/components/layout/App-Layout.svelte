<script lang="ts">
    import * as Sidebar from "$lib/components/ui/sidebar/index.js";
    import AppSidebar from "$lib/components/layout/Sidebar-Layout.svelte";
    import Navbar from "$lib/components/layout/Navbar-Layout.svelte";

    let { children } = $props();
</script>

<Sidebar.Provider class="min-h-full">
    <AppSidebar />
    <main class="flex-1 block">
        <Navbar />
        {@render children?.()}
    </main>
</Sidebar.Provider>
