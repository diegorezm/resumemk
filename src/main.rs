mod cli;
mod resume_builder;

use cli::{get_output_path, init_cli, Command, OutputType};
use resume_builder::ResumeBuilder;

fn main() {
    let cli = init_cli();
    match cli.command {
        Command::Build {
            input,
            output,
            output_type,
            stylesheet,
            title,
        } => {
            let output_path = get_output_path(&input, &output, &output_type);
            let mut document_title = output_path.to_str().unwrap().to_string();
            let mut resume_builder = ResumeBuilder::new();

            if let Some(stylesheet) = stylesheet {
                resume_builder.set_stylesheet(stylesheet);
            }

            if let Some(title) = title {
                document_title = title;
            }

            let pdf_options = headless_chrome::types::PrintToPdfOptions::default();

            match output_type {
                OutputType::HTML => {
                    match resume_builder.save_to_html(
                        input.as_path(),
                        output_path.as_path(),
                        document_title.as_str(),
                    ) {
                        Ok(path) => {
                            println!("Successfully saved to {}", path);
                        }
                        Err(e) => {
                            println!("Error: {}", e);
                        }
                    }
                }
                OutputType::PDF => {
                    match resume_builder.save_to_pdf(
                        input.as_path(),
                        output_path.as_path(),
                        pdf_options,
                        document_title.as_str(),
                    ) {
                        Ok(_) => {
                            println!("Successfully saved to {}", output_path.to_str().unwrap());
                        }
                        Err(e) => {
                            println!("Error: {}", e);
                        }
                    }
                }
            }
        }
    }
}
