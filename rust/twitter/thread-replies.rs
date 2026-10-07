// How do I get replies or a full thread on X (Twitter)?
// https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x
// Run: cargo run --bin twitter-thread-replies

use openhandle::Client;

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let mut replies = client
        .twitter()
        .post("940100000000000001")
        .comments()
        .items(None);
    while let Some(reply) = replies.next().await {
        let reply = reply?;
        let author = reply.author.and_then(|author| author.handle);
        println!("{:?} {:?}", author, reply.text);
    }
    Ok(())
}
