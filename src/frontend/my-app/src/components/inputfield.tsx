import { Input } from '@/components/ui/input';
import React from 'react';

interface InputFieldProps {
    placeholder?: string;
}

const InputField: React.FC<InputFieldProps> = ({ placeholder }) => {
    return (
        <div className='p-5'>
            <Input className='font-raleway text-neutral-400' placeholder={placeholder} />
        </div>
    );
};

export default InputField;
