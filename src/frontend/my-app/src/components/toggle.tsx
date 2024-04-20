import { Switch } from "@/components/ui/switch"

const Toggle = () => {
    return (
        <div className="flex flex-row gap-3 pb-5">
            <p className="text-neutral-100 text-lg"><b>IDS</b></p>
            <Switch/>
            <p className="text-neutral-100 text-lg"><b>BFS</b></p>
        </div>
    );
};

export default Toggle;