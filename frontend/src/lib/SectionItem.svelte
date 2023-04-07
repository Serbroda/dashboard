<script lang="ts">
    export let url: string;
    export let name: string;
    let status: "online" | "offline" = "offline";
    let checked = false;

    function chk(target, times, delay) {
        return new Promise((res, rej) => {                       // return a promise

            (function rec(i) {                                   // recursive IIFE
                fetch(target, {mode: 'no-cors'}).then((r) => {   // fetch the resourse
                    res(r);                                      // resolve promise if success
                }).catch( err => {
                    if (times === 0)                             // if number of tries reached
                        return rej(err);                         // don't try again

                    setTimeout(() => rec(--times), delay )       // otherwise, wait and try
                });                                              // again until no more tries
            })(times);

        });
    }

    chk(url, 3, 1000).then( image => {
        status = "online"
    }).catch( err => {
        status = "offline"
    });
</script>

<a target="_blank" style="text-decoration: none" class="flex bg-surface-700 p-2 rounded-md"
   href={url}>{name}
    <div class="flex-1"></div>
    <div class={`w-2 self-start aspect-square rounded-full animate-pulse ${status === "online" ? "bg-green-500" : "bg-red-500"}`}></div>
</a>