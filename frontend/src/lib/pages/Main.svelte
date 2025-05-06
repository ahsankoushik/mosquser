<script lang="ts">
    import { AtSign, Database, LogOut } from "@lucide/svelte";
import Button from "../components/ui/button/button.svelte";


</script>

<div class="w-screen h-screen flex">
    <div class="basis-1/12  ">
        <div class="h-full flex flex-col justify-between">
            <div class=" grid py-4 gap-2">
                <div class="flex justify-center">
                    <Button
                        class="bg-white  hover:bg-gray-200"
                    >
                        <Database color={"black"}/>
                    </Button>
                </div>
                <div class="flex justify-center">
                    <Button
                        class="bg-white  hover:bg-gray-200"
                    >
                        <AtSign color={"black"}/>
                    </Button>
                </div>
            </div>
    
            <div
                class="mb-6 flex justify-center"
            >
                <Button>
                    <LogOut />
                </Button>
            </div>
        </div>
        
       
    </div>
    
    <div class="basis-3/12 border">
    
    </div>
    <div class="basis-8/12 border">
    
    </div>
</div>