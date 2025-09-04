<script lang="ts">
    import ArrowLeft from "@lucide/svelte/icons/arrow-left";
    import EllipsisVertical from "@lucide/svelte/icons/ellipsis-vertical";
    import profile10 from "$lib/assets/avatars/10.png";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
</script>

<header class="sticky top-0 z-50 min-w-full">
    <nav class="relative bg-sidebar shadow-sm py-4 flex justify-between align-middle items-center">
        <div class="ml-5 flex align-middle items-center space-x-4">
            <ArrowLeft />
            <img src={profile10} alt="" class="w-[47px] h-[47px]" />
            <span class="text-lg">Angga</span>
        </div>

        <div class="pr-3">
            <DropdownMenu.Root>
            <DropdownMenu.Trigger><EllipsisVertical /></DropdownMenu.Trigger>
                <DropdownMenu.Content>
                <DropdownMenu.Group>
                    <DropdownMenu.Item>Profile</DropdownMenu.Item>
                </DropdownMenu.Group>
                </DropdownMenu.Content>
            </DropdownMenu.Root>
        </div>
    </nav>
</header>
