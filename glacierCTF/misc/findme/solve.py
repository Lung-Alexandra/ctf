import pikepdf

pdf_path = 'chall.pdf'
output_png_path = "file.png"

with pikepdf.open(pdf_path) as pdf:
    combined_data = b''  # all combined stream data

    for obj_id, obj in enumerate(pdf.objects):
        if isinstance(obj, pikepdf.Stream):
            length = obj.get("/Length")
            if length == 100 or length == 90:
                #  concatenate streams in order to obtain final result
                data = obj.read_bytes()
                combined_data += data

    with open(output_png_path, 'wb') as png_file:
        png_file.write(combined_data)