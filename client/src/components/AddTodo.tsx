import React from 'react'
import {useState}from "react"
import {useForm} from '@mantine/form';
import {Button, Group, Modal, TextInput, Textarea} from "@mantine/core"

function AddTodo() {
    const [open, setOpen] = useState(false)

    const form = useForm({
        initialValues : {
            title : "",
            body : "",
        }
    });

    function createTodo () {

    }

    return (
        <>
        <Modal opened={open} onClose={() => setOpen(false)} title="Create todo"> Test </Modal>
    <form onSubmit={form.onSubmit(createTodo)}>
        <TextInput required mb={12} label="Todo" placeholder="What do you want to do?" {...form.getInputProps("title")} />
        <Textarea />

        <Button type="submit">Create todo</Button>
    </form>
        <Group position="center">
            <Button fullWidth mb={12} onClick={() => setOpen(true)}> Add Todo </Button>
        </Group>
        </>
    )
}

export default AddTodo