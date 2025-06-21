<script lang="ts">
    const {child, helpText, topic} = $props();
</script>


<div class="group relative">
        {@render child(topic)}
        <p
            class="absolute bottom-[-12] hidden group-hover:block w-auto bg-gray-600 text-white text-xs z-10 p-0.5 rounded-md"
        >
            {helpText}
        </p>
    </div>

